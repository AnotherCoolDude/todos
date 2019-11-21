<template>
    <li>
        <div class="todoItemContainer">
            <h3 class="item_row">{{ todo.title }}</h3>
            <p class="item_row">Zeitraum</p>
            <Datepicker :value="todo.startDate" class="item_left" ></Datepicker>
            <Datepicker :value="todo.endDate" class="item_right" ></Datepicker>
            <p class="item_row">Zuständiger</p>
            <Autocomplete class="item_row" :search="search" :get-result-value="getResultValue"></Autocomplete>
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
            },
           

        },
        components: {
            Datepicker,
            Autocomplete,
            VueTimepicker
        },
        data: function () {
            return {
                testEmployees: [{
                    name: "Alex"
                }, {
                    name: "Peter"
                }, {
                    name: "Thomas"
                }]
            }
        },
        methods: {
            search: function (input) {
                if (input.length < 1) {
                    return []
                }
                return this.testEmployees.filter(employee => {
                    return employee.name.toLowerCase()
                        .startsWith(input.toLowerCase())
                })
            },
            getResultValue: function (result) {
                return result.name
            },
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