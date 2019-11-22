<template>
    <div>
        <section class="transferapp" v-cloak>
            <header class="header">
                <h1>Basecamp to Proad</h1>
                <h3 v-if="errorMessage.length > 0">{{errorMessage}}</h3>
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

// import Wails from "@wailsapp/runtime";

export default {
    name: "app",
    data: function() {
        return {
            errorMessage: "",
            loading: false,
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
            }],
            projects: []
        };
    },
    mounted() {
        if (!this.loading) {
            this.loading = true;
            this.loadProjects();
        }
    },
    components: {
        pitem,
        titem
    },
    methods: {
        loadProjects: function() {
            window.backend.Basecamp.Login()
                .then(() => {
                    window.backend.Basecamp.FetchProjects()
                        .then(() => {
                            window.backend.Basecamp.GetProjects().then(pp => {
                                this.projects = pp.filter(p => {
                                    p.nr != ""
                                });
                                this.loading = false;
                            });
                        })
                        .catch(error => {
                            this.errorMessage = error;
                        });
                })
                .catch(error => {
                    this.errorMessage = error;
                });
        }
    }      
}
</script>

<style>

</style>