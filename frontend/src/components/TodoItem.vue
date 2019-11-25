<template>
    <li>
        <div class="todoItemContainer">
            <h3 class="item_title">{{ todo.title }}</h3>
            <p class="item_desc item_col_left">Zeitraum</p>
            <Datepicker v-model="todo.startDate" class="item_input_top item_col_left item_input" ></Datepicker>
            <Datepicker v-model="todo.endDate" class="item_input_bottom item_col_left item_input" ></Datepicker>
            <p class="item_desc item_col_mid">Zuständiger</p>
            <Autocomplete base-class="search_emp" class="item_input_top item_col_mid item_input" :search="search" :get-result-value="getResultValue" @submit="submit"></Autocomplete>
            <p class="item_desc item_col_right">Erledigt / Gesamtaufwand</p>
            <input class="item_input_top item_col_right item_input" v-model.number="todo.workAmountDone" type="number" step="any">
            <input class="item_input_bottom item_col_right item_input" v-model.number="todo.workAmountTotal" type="number" step="any">
        </div>
    </li>
</template>

<script>
    import Datepicker from 'vuejs-datepicker';
    import Autocomplete from '@trevoreyre/autocomplete-vue'

    export default {
        props: {
            todo: {
                title: String,
                startDate: Date,
                endDate: Date,
                workAmountTotal: Number,
                workAmountDone: Number,
                projectnr: String,
                assignee: {firstname: String, lastname: String, urno: Number}
            },
            employees: Array
        },
        components: {
            Datepicker,
            Autocomplete
        },
        data: function () {
            return {
            }
        },
        methods: {
            search: function (input) {
                if (input.length < 1) {
                    return []
                }
                return this.employees.filter(employee => {
                    const s = employee.name.split(" ")
                    var match = false
                    for (const n of s) {
                        if (n.toLowerCase()
                            .startsWith(input.toLowerCase())) {
                                match = true
                            }
                    }
                    return match
                })
            },
            getResultValue: function (result) {
                return result.name
            },
            submit: function(result) {
                const split = result.name.split(" ")
                this.todo.assignee = {firstname: split[0], lastname: split[1], urno: result.urno}
                this.$emit("todoUpdated", this.todo, this.$vnode.key)
            }
        }
    }
</script>

<style>

.item_input {
    width: 100px;
}

.search_emp ul {
    background: white;
    width: 80%;
}

.search_emp-result-list li {
    background: white;
    text-align: center;
    margin: 0px;
    padding: 0px;
    border: none;
}

.search_emp-result-list li:hover {
    background: rgb(211, 222, 255);
}

.todoItemContainer {
    display: grid;
    grid-template-columns: auto;
    grid-template-rows: auto;
}

.todoItemContainer .item_title {
    grid-row: 1;
    grid-column: 1 / end;
}

.todoItemContainer .item_desc {
    grid-row: 2;
}

.todoItemContainer .item_input_top {
    width: 100%;
    grid-row: 3;
}

.todoItemContainer .item_input_bottom {
    width: 100%;
    grid-row: 4;
}

.todoItemContainer .item_col_left {
    grid-column: 1;
}

.todoItemContainer .item_col_mid {
    grid-column: 2;
}

.todoItemContainer .item_col_right {
    grid-column: 3;
}

</style>