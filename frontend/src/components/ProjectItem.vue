<template>
  <li class="item" @click="selected">
      <div class="container">
          <p class="jobnr">{{ project.nr }}</p>
          <p class="counter">{{ project.todos.length }} Todos</p>
      </div>
  </li>
</template>

<script>
export default {
    props: {
        project: {nr: String, todos: []}
    },
    methods: {
        selected: function() {
            this.$emit("projectSelected", this.$vnode.key)
        }
    }
}
</script>

<style scoped >

.item {
    min-height: 20px;
    width: 100%;
    padding: 0;
}

.container {
    display: grid;
    grid-template-columns: 100%;
    grid-template-rows: 50% 50%;
    margin: 8px;
}

.jobnr {
    grid-column: 1;
    grid-row: 1;
}

.counter {
    grid-column: 1;
    grid-row: 2;
    font-size: 12px;
    opacity: 0.8;
}

</style>