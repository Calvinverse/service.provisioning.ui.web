<template>
  <div class="home">
    <EnvironmentsList />
  </div>
</template>

<script>
// @ is an alias to /src
import EnvironmentsList from '@/components/EnvironmentsList.vue'

export default {
  name: 'Home',
  components: {
    EnvironmentsList
  }
}
</script>
