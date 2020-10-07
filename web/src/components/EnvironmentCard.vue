<template>

</template>

<script lang="ts">
import Vue from 'vue'
import {
  namespace,
  State
} from 'vuex-class'
import Component from 'vue-class-component'

const environmentModule = namespace('environment')

@Component
export default class UserCard extends Vue {
  @State('profile') profile!: ProfileState
  @environmentModule.Action('login') login: any
  @environmentModule.Action('logout') logout: any
  @environmentModule.Getter('fullName') fullName!: string
  @environmentModule.Getter('gravatarImage') gravatarImage!: string
}
</script>
