<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |      
      <router-link to="/q1">Word Counter</router-link> |
      <router-link to="/q2">Excel Column</router-link> |
      <router-link to="/q3">CRUD</router-link> |
      <router-link to="/q4">Sample Service</router-link> |
      <router-link to="/q5">Goroutine</router-link>

    </div>
    <router-view/>
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin: 20px 20px;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
