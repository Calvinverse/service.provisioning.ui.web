<template>
  <v-container>
    <!-- Anonymous -->
    <v-menu
      bottom
      offset-y
      v-if="!profile.isAuthenticated"
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
      bottom
      offset-y
      v-if="profile.isAuthenticated"
      >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          color="primary"
          v-bind="attrs"
          v-on="on"
        >
          <strong>{{ fullName }}</strong>
        </v-btn>
      </template>
      <v-list v-if="profile.isAuthenticated">
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
import Vue from 'vue'
import {
  namespace,
  State
} from 'vuex-class'
import Component from 'vue-class-component'
import { ProfileState } from '../store/profile/ProfileState'
import { profile } from '../store/profile'

const profileModule = namespace('profile')

@Component
export default class UserLoginState extends Vue {
  @State('profile') profile!: ProfileState
  @profileModule.Action('login') login: any
  @profileModule.Action('logout') logout: any
  @profileModule.Getter('fullName') fullName!: string

  mounted () {
    this.setUser()
  }

  setUser () {
    // this.user = this.$AuthService.getUser()
  }
}
</script>
