<template>
    <div>
        <section class="transferapp" v-cloak>
            <header class="header">
                <h1>Basecamp to Proad</h1>
                <h3 v-if="errorMessage.length > 0">{{errorMessage}}</h3>
            </header>
            <section class="projectlist">
                <ul>
                    <li is="pitem" 
                    v-for="(project, index) in projects" 
                    v-bind:project="project" 
                    v-bind:key="index"
                    @projectSelected="showTodos"></li>
                </ul>
            </section>
            <section class="todoconfig">
                <ul>
                    <li is="titem"
                        v-for="(todo, index) in selectedProject.todos"
                        v-bind:key="index"
                        v-bind:todo="todo"></li>
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
            projects: [],
            selectedProject: Object,
            employees: []
        };
    },
    mounted() {
        if (!this.loading && this.projects.length == 0) {
            this.loading = true;
            this.login();
            this.getProjects();
        }
    },
    components: {
        pitem,
        titem
    },
    methods: {
        login: function() {
            window.backend.Basecamp.Login()
                .catch(error => {
                    this.errorMessage = error;
                });
        },
        getProjects: function() {
            window.backend.Basecamp.FetchProjects()
                .then(() => {
                    window.backend.Basecamp.GetProjects().then(pp => {
                        this.projects = pp.filter(p => {
                            return p.nr != "";
                        });
                        this.selectedProject = this.projects[0];
                    })
                    this.loading = false;
                })
                .catch(error => {
                    this.errorMessage = error;
                });
        },
        showTodos: function(key) {
            this.selectedProject = this.projects[key];
        }
    }      
}
</script>

<style>

</style>