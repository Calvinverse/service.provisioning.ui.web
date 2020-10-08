<template>

</template>

<script lang="ts">
import Vue from 'vue'
import {
  namespace,
  State
} from 'vuex-class'
import Component from 'vue-class-component'
import { Environment } from '../store/environment'

const environmentModule = namespace('environment')

@Component
export default class EnvironmentCard extends Vue {
  @State('environment') environment!: Environment
  @environmentModule.Action('get') get: any
  @environmentModule.Action('getAll') getAll: any
  @environmentModule.Action('delete') delete: any
  // @environmentModule.Getter('fullName') fullName!: string
  // @environmentModule.Getter('gravatarImage') gravatarImage!: string
}
</script>
