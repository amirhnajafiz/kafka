import { defineStore } from 'pinia'

import data from '../locals/data.json'

export const useInformationStore = defineStore({
  id: 'information',
  state: () => ({
    data: data
  }),
  getters: {
    getFirstName: (state) => state.data.firstname,
    getLastName: (state) => state.data.lastname,
    getAge: (state) => state.data.age,
  },
})
