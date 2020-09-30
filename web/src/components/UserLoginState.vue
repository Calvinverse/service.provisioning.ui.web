<template>
  <v-container>
    <!-- Anonymous -->
    <v-menu
      top
      :offset-y="offset"
      v-if="!user"
      >

      <template v-slot:activator="{ on, attrs }" >
        <v-btn
          color="primary"
          v-bind="attrs"
          v-on="on"
        >
          <strong>Anonymous</strong>
        </v-btn>
      </template>
      <v-list>
        <v-list-item>
          <v-list-item-title @click="login"><i class="fa fa-lock" /> Login</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <!-- Authenticated -->
    <v-menu
      top
      :offset-y="offset"
      v-if="user"
      >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          color="primary"
          v-bind="attrs"
          v-on="on"
        >
          <strong>{{ user.name }}</strong>
        </v-btn>
      </template>
      <v-list v-if="user">
        <v-list-item>
          <v-list-item-title>
            <strong>Account</strong>
          </v-list-item-title>
          <v-list-item-title>
            <i class="fa fa-user" /> Profile
          </v-list-item-title>
          <v-list-item-title>
            <i class="fa fa-wrench" /> Settings
          </v-list-item-title>
          <v-list-item-title
            @click="logout"
          >
            <i class="fa fa-lock" /> Logout
          </v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component
export default class UserLoginState extends Vue {
  user: null

  mounted() {
    this.setUser()
  }

  logout() {
    this.user = null
    this.$AuthService.logout()
  }

  login() {
    // this.$AuthService.loginPopup() //with a popup
    this.$AuthService.loginRedirect() //with a redirect
  }

  setUser() {
    this.user = this.$AuthService.getUser()
  }
}
</script>
