import { Module } from 'vuex'
import { getters } from './getters'
import { actions } from './actions'
import { mutations } from './mutations'
import {
  EnvironmentState,
} from './types'

import { RootState } from '../RootState'

export const state: EnvironmentState = {
  error: false,

  id: '',
  name: '',
  description: '',
  createdOn: new Date(),
  destroyBy: new Date(),
  resources: new Array<string>(),
  status: '',
  tags: new Array<string>(),
  version: ''
}

const namespaced = true

export const environment: Module<EnvironmentState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations
}
