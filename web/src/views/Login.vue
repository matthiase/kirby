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
                    Sign into your account
                  </h2>
                  <p>
                    Or
                    <router-link to="/register">
                      register for a free account
                    </router-link>
                  </p>
                </div>

                <form @submit.prevent="handleLogin" class="pt-5">
                  <b-field label="Email address">
                    <b-input type="email" v-model="user.email" />
                  </b-field>
                  <div class="field pt-2">
                    <div class="is-clearfix pb-2">
                      <div class="is-pulled-left">
                        <label class="label">Password</label>
                      </div>
                      <div class="is-pulled-right font-sm">
                        <a href="#">Forgot password?</a>
                      </div>
                    </div>
                    <b-input type="password" v-model="user.password" password-reveal />
                  </div>
                  <div class="has-text-centered py-5">
                    <b-checkbox size="is-small">Remember this device</b-checkbox>
                  </div>
                  <div class="mt-3 mb-5">
                    <b-button type="is-info" native-type="submit" expanded>
                      Sign in
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
import { mapActions } from 'vuex'
import router from "@/router"

export default {
  name: "LoginView",
  data() {
    return {
      user: {
        email: "",
        password: ""
      }
    }
  },
  methods: {
    ...mapActions("authentication", ["login"]),
    async handleLogin() {
      const currentUser = await this.login(this.user)
      if (currentUser) {
        router.push("/profile")
      }
    }
  }
}
</script>
