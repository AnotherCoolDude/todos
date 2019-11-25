<template>
    <li>
        <div class="todoItemContainer">
            <h3 class="item_row">{{ todo.title }}</h3>
            <p class="item_row">Zeitraum</p>
            <Datepicker :value="todo.startDate" class="item_left" ></Datepicker>
            <Datepicker :value="todo.endDate" class="item_right" ></Datepicker>
            <p class="item_row">Zuständiger</p>
            <Autocomplete class="item_row" :search="search" :get-result-value="getResultValue" @submit="submit"></Autocomplete>
            <p class="item_row">Erledingt / Gesamtaufwand</p>
            <VueTimepicker class="item_left" format="HH:mm" v-model="todo.workAmountDone"></VueTimepicker>
            <VueTimepicker class="item_right" format="HH:mm" v-model="todo.workAmountTotal"></VueTimepicker>
        </div>
    </li>
</template>

<script>
    import Datepicker from 'vuejs-datepicker';
    import Autocomplete from '@trevoreyre/autocomplete-vue'
    import VueTimepicker from 'vue2-timepicker/src/vue-timepicker.vue';

    export default {
        props: {
            todo: {
                title: String,
                startDate: Date,
                endDate: Date,
                workAmountTotal: {HH: String, mm: String},
                workAmountDone: {HH: String, mm: String},
                projectnr: String,
                assignee: {firstname: String, lastname: String, urno: Number}
            },
            employees: Array
        },
        components: {
            Datepicker,
            Autocomplete,
            VueTimepicker
        },
        data: function () {
            return {
                ee: this.employees
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
.todoItemContainer {
    display: grid;
    grid-template-columns: 50% 50%;
    grid-template-rows: 100/7%;
}

.todoItemContainer .item_row {
    grid-column: 1 / end;
}

.todoItemContainer .item_left {
    grid-column: 1;
}

.todoItemContainer .item_right {
    grid-column: 2;
}

</style>