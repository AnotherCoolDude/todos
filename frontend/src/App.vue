<template>
    <div>
        <section class="transferapp" v-cloak>
            <header class="header">
                <h1>Basecamp to Proad</h1>
            </header>
            <section class="projectlist">
                <ul>
                    <li is="pitem" v-bind:job="Testjob"></li>
                    <li is="pitem" v-bind:job="Testjob"></li>
                </ul>
            </section>
            <section class="todoconfig">
                <ul>
                    <li is="titem" v-bind:todo="TestTodo"></li>
                    <li is="titem" v-bind:todo="TestTodo"></li>
                    <li is="titem" v-bind:todo="TestTodo"></li>
                </ul> 
            </section>
            <section class="footer">
                    <button>Tranfer</button>
            </section> 
        </section>
    </div>
</template>

<script>
import "./assets/css/transferapp.css"

import pitem from "./components/ProjectItem.vue";
import titem from "./components/TodoItem.vue";

export default {
    name: "app",
    data: function() {
        return {
            Testjob: {
                nr: "SEIN-0001-0001",
                amount: 5
                },
            TestTodo: {                
                title: "Druck",
                startDate: new Date(),
                endDate: new Date(),
                workAmountTotal: {HH: "10", mm: "00"},
                workAmountDone: {HH: "02", mm: "30"}
                },
            Todos: [{
                title: String,
                startDate: Date,
                endDate: Date,
                workAmountTotal: {HH: String, mm: String},
                workAmountDone: {HH: String, mm: String}
            }]
        };
    },
    components: {
        pitem,
        titem
    },
    methods: {
        loadProjects: function() {
            console.log("wip");
        }
    }      
}
</script>

<style>

</style>