<template>
  <b-navbar transparent>
    <template slot="brand">
      <b-navbar-item tag="router-link" :to="{ path: '/' }">
        <img src="../assets/logo.svg" alt="Kirby Logo" class="image is-32x32" />
      </b-navbar-item>
    </template>
    <template slot="end">
      <b-navbar-item tag="div">
        <b-navbar-dropdown v-if="authenticated">
          <template slot="label">
            <b-image src="https://avatars1.githubusercontent.com/u/210818?s=60&v=4" class="is-24x24" rounded />
          </template>
          <b-navbar-item tag="router-link" to="/profile">
            Your profile
          </b-navbar-item>
          <b-navbar-item href="#" @click="handleSignOut">
            Sign out
          </b-navbar-item>
        </b-navbar-dropdown>
        <div v-else class="buttons">
          <b-button type="is-info" tag="router-link" to="/login">
            Sign in
          </b-button>
        </div>
      </b-navbar-item>
    </template>
  </b-navbar>
</template>

<script>
import { mapState, mapActions } from "vuex"
import router from "@/router"

export default {
  name: "Navbar",
  computed: {
    ...mapState("authentication", ["authenticated", "user"])
  },
  methods: {
    ...mapActions("authentication", ["logout"]),
    handleSignOut() {
      this.logout()
      router.push("/")
    }
  }
}
</script>
