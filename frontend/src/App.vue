<template>
  <v-app>
      <v-app-bar
          height="50"
          fixed=true>
          yoooz
          <div class="links">
              <router-link to="/">Home</router-link> |
              <router-link to="/about">About</router-link>
          </div>
      </v-app-bar>
      <v-container>
          <router-view/>
      </v-container>
  </v-app>
</template>

<style lang="scss">
 #app {
     font-family: Avenir, Helvetica, Arial, sans-serif;
     -webkit-font-smoothing: antialiased;
     -moz-osx-font-smoothing: grayscale;
     text-align: center;
     color: #2c3e50;
 }
 
 header.v-toolbar div.v-toolbar__content {
     justify-content: space-between;
     div.links {
         right: 0;
     }

    a {
        font-weight: bold;
        color: #2c3e50;
        
        &.router-link-exact-active {
            color: #42b983;
        }
    }
 }

 div.container {
     margin-top: 50px;
 }
</style>
