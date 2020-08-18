<template>
  <div id="app">
    <div v-if="alert.message">
      <b-notification
        :type="`is-${alert.type} is-light`"
        @close="handleNotificationClosed"
        aria-close-label="Close notification"
        role="alert"
      >
        {{ alert.message }}
      </b-notification>
    </div>
    <router-view />
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex"

export default {
  computed: {
    ...mapState({
      alert: state => state.alert
    })
  },
  methods: {
    ...mapActions({
      clearAlert: "alert/clear"
    }),
    handleNotificationClosed() {
      this.clearAlert()
    }
  },
  watch: {
    $route() {
      // clear alert on location change
      this.clearAlert()
    }
  }
}
</script>
