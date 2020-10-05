<template>
  <v-menu
      bottom
      origin="center center"
      transition="scale-transition"
      color="secondary"
      :rounded="b-xl"
      :offset-y="offset"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-card
          color="secondary"
          v-bind="attrs"
          v-on="on"
          height="100%"
          >
          <v-container
            fill-height>
            <v-row
              >
              <v-col
                cols="2">
                <v-avatar
                  v-if="!profile.isAuthenticated">
                  <v-icon x-large>mdi-account-circle</v-icon>
                </v-avatar>
                <v-avatar
                  v-if="profile.isAuthenticated">
                  <v-img
                    :src="gravatarImage"
                  />
                </v-avatar>
              </v-col>
              <v-col>
                <v-card-title>{{ userName }}</v-card-title>
              </v-col>
            </v-row>
          </v-container>
        </v-card>
      </template>
      <v-list>
        <v-list-item  @click="userChangeAction">
          <v-list-item-icon>
            <v-icon v-text="userChangeActionIcon"></v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title v-text="userChangeActionText"></v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
  </v-menu>
</template>

<script lang="ts">
import Vue from 'vue'
import {
  namespace,
  State
} from 'vuex-class'
import Component from 'vue-class-component'
import { ProfileState } from '../store/profile/ProfileState'

const profileModule = namespace('profile')

@Component
export default class UserCard extends Vue {
  @State('profile') profile!: ProfileState
  @profileModule.Action('login') login: any
  @profileModule.Action('logout') logout: any
  @profileModule.Getter('fullName') fullName!: string
  @profileModule.Getter('gravatarImage') gravatarImage!: string

  get userName () {
    return this.profile.isAuthenticated ? this.fullName : 'Anonymous'
  }

  get userChangeActionIcon () {
    return this.profile.isAuthenticated ? 'mdi-logout' : 'mdi-login'
  }

  get userChangeActionText () {
    return this.profile.isAuthenticated ? 'Log out' : 'Log in'
  }

  get userChangeAction () {
    return this.profile.isAuthenticated ? this.logout : this.login
  }
}
</script>
