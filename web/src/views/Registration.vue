<template>
  <div class="min-h-screen has-background-white-bis">
    <div class="section">
      <div class="columns is-mobile is-centered">
        <div class="column is-one-third" style="min-width:24em; max-width:32em;">
          <div>
            <div class="box">
              <div class="px-5">
                <div class="has-text-centered py-3">
                  <router-link to="/">
                    <figure class="image is-32x32 is-inline-block">
                      <img src="../assets/logo.svg" alt="Kirby Logo" />
                    </figure>
                  </router-link>
                  <h2 class="is-size-4 has-text-weight-bold">
                    Register for a new account
                  </h2>
                  <p>
                    Or
                    <router-link to="/login">
                      login with your existing account
                    </router-link>
                  </p>
                </div>
                <form @submit.prevent="handleSubmit" class="pt-5">
                  <b-field label="Full name">
                    <b-input type="text" v-model="user.name" />
                  </b-field>
                  <b-field label="Email address">
                    <b-input type="email" v-model="user.email" />
                  </b-field>
                  <b-field label="Password">
                    <b-input type="password" v-model="user.password" autocomplete="new-password" password-reveal />
                  </b-field>
                  <div class="py-5 has-text-centered has-text-grey">
                    By clicking the "Register" I agree to the terms of service
                  </div>
                  <div class="mt-3 mb-5">
                    <b-button type="is-primary" native-type="submit" expanded :loading="loading">
                      Register
                    </b-button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex"
import router from "@/router"

export default {
  name: "RegistrationView",
  data() {
    return {
      user: {
        name: "",
        email: "",
        password: ""
      }
    }
  },
  computed: {
    ...mapState("authentication", ["authenticated", "loading"])
  },
  methods: {
    ...mapActions("authentication", ["register"]),
    async handleSubmit() {
      const currentUser = await this.register(this.user)
      if (currentUser) {
        router.push("/profile")
      }
    }
  }
}
</script>
