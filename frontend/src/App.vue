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
                    <li>
                        <div>
                        <h3>Layout</h3>
                        <p>Datum</p>
                        <p>Zuständiger</p>
                        <p>Aufwand</p>
                        </div>
                    </li>
                    <li>
                        <div>
                        <h3>Layout</h3>
                        <p>Datum</p>
                        <p>Zuständiger</p>
                        <p>Aufwand</p>
                        </div>
                    </li>
                    <li>
                        <div>
                        <h3>Layout</h3>
                        <p>Datum</p>
                        <p>Zuständiger</p>
                        <p>Aufwand</p>
                        </div>
                    </li>
                </ul>
                <button>Tranfer</button>
            </section>
        </section>
    </div>
</template>

<script>
import "./assets/css/transferapp.css"

import pitem from "./components/ProjectItem.vue";

export default {
    name: "app",
    data: function() {
        return {
            Testjob: {nr: "SEIN-0001-0001", amount: "5"}
        };
    },
    components: {
        pitem
    }
}
</script>

<style>

</style>