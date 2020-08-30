<template>
  <section v-if="currentUser">

    <b-modal v-model="isImageUploadActive" has-modal-card full-screen>
      <template>
        <image-upload @close="handleImageUploadCancel" title="Change avatar"  :imageUrl="avatarUrl" />
      </template>
    </b-modal>

    <div class="profile-card has-text-centered">
      <a href="javascript:void(0);" @click="handleAvatarClick">
        <vue-avatar
          :username="currentUser.name"
          :size="150"
          class="avatar-image"
          backgroundColor="rgb(6, 148, 162)"
          color="#EFEFEF"
          :customStyle="avatarStyle"
        />
      </a>
      <div class="mt-5 px-3">
        <form v-if="isEditing" @submit.prevent="handleSaveProfileClick" class="px-2">
          <b-field>
            <b-input v-model="user.name" placeholder="Full name" />
          </b-field>
          <b-field>
            <b-input type="email" v-model="user.email" placeholder="Email address" />
          </b-field>
          <div class="columns mt-4">
            <div class="column">
              <b-button type="is-primary" native-type="submit" expanded>
                Save
              </b-button>
            </div>
            <div class="column" @click="handleCancelProfileClick">
              <b-button expanded>
                Cancel
              </b-button>
            </div>
          </div>
        </form>
        <div v-else>
          <h5 class="is-size-5">{{ currentUser.name }}</h5>
          <h6 class="is-size-6 has-text-grey">{{ currentUser.email }}</h6>
          <b-button expanded class="my-5" @click="handleEditProfileClick">
            Edit Profile
          </b-button>
        </div>
      </div>
    </div>

    <b-tabs position="is-centered" class="block is-clearfix">
      <b-tab-item label="Dashboards"></b-tab-item>
      <b-tab-item label="Collections"></b-tab-item>
      <b-tab-item label="Preferences"></b-tab-item>
    </b-tabs>
    <div class="columns">
      <div class="column is-one-quarter"></div>
      <div class="column">
        &nbsp;
      </div>
    </div>
  </section>
</template>

<style lang="scss" scoped>
.profile-card {
  @media (min-width: 768px) {
    float: left;
    width: 240px;
    z-index: 100;
  }
  .avatar-image {
    margin: 0 20%;
  }
}
</style>

<script>
import { mapActions, mapState } from "vuex"
import VueAvatar from "vue-avatar"
import ImageUpload from "@/components/ImageUpload"

export default {
  name: "ProfileView",
  components: {
    VueAvatar,
    ImageUpload
  },
  data() {
    return {
      isEditing: false,
      isImageUploadActive: false,
      avatarUrl: 'https://images.unsplash.com/photo-1519119012096-c145def61801?ixlib=rb-1.2.1&auto=format&fit=crop&w=1280&q=80',
      user: {
        name: "",
        email: ""
      },
      avatarStyle: {
        display: "inline-block",
        fontFamily: "inherit"
      }
    }
  },
  computed: {
    ...mapState("authentication", ["currentUser"])
  },
  methods: {
    ...mapActions("authentication", ["updateProfile"]),
    handleEditProfileClick() {
      this.user.name = this.currentUser.name
      this.user.email = this.currentUser.email
      this.isEditing = true
    },
    handleCancelProfileClick() {
      this.isEditing = false
    },
    async handleSaveProfileClick() {
      await this.updateProfile(this.user)
      this.isEditing = false
    },
    handleAvatarClick() {
      this.isImageUploadActive = true
    },
    handleImageUploadCancel() {
      this.isImageUploadActive = false
    }
  }
}
</script>
