<template>
  <v-container>
    <li v-for="(item, index) in environmentList" :key="item.id" :data-index="index">
      <v-card class="event-card">
        <v-layout row>
          <v-layout column justify-space-between style="padding: 0.8em 1.3em; max-width: 390px;">
            <div>
              <h1 class="name"></h1>
              <h3 class="date"></h3>
            </div>
            <div>
              <p class="desc"></p>
              <div class="location">
                <v-icon v-if="item.address">location_on</v-icon>

              </div>
            </div>
            <div class="date-ribbon">
              <h2></h2>
              <h1></h1>
            </div>
          </v-layout>
        </v-layout>
      </v-card>
    </li>
    <v-bottom-navigation>
      <v-btn>
        <span>Add</span>
        <v-icon>mdi-plus-thick</v-icon>
      </v-btn>

      <v-btn>
        <span>Delete</span>
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </v-bottom-navigation>
  </v-container>
</template>

<script lang="ts">

import Vue from 'vue'
import {
  namespace,
  State
} from 'vuex-class'
import Component from 'vue-class-component'
import { Environments } from '../store/environment'

const environmentModule = namespace('environment')

@Component
export default class EnvironmentsList extends Vue {
  @State('environment') environments!: Environments
  @environmentModule.Action('get') get: any
  @environmentModule.Action('getAll') getAll: any
  @environmentModule.Action('delete') delete: any
  // @environmentModule.Getter('fullName') fullName!: string
  // @environmentModule.Getter('gravatarImage') gravatarImage!: string
}
</script>
