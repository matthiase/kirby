<template>
  <div class="modal-card" style="width: auto">
    <header class="modal-card-head">
      <p class="modal-card-title">{{ title }}</p>
      <button type="button" class="delete" @click="$emit('close')" />
    </header>
    <section class="modal-card-body">
      <cropper
        :src="imageUrl"
        :stencil-props="{
          aspectRatio: 1
        }"
        ref="cropper"
        class="cropper"
      />
    </section>
    <footer class="modal-card-foot" style="justify-content:space-between; align-items:center;">
      <div>
        <b-upload v-model="file">
          <a class="button">
            <b-icon icon="upload"></b-icon>
            <span>Upload</span>
          </a>
        </b-upload>
      </div>
      <div>
        <button class="button" type="button" @click="$emit('close')">Cancel</button>
        <button class="button is-primary">Save</button>
      </div>
    </footer>
  </div>
</template>

<script>
import { Cropper } from "vue-advanced-cropper"
export default {
  name: "ImageUpload",
  components: {
    Cropper
  },
  props: {
    imageUrl: {
      type: String
    },
    title: {
      type: String,
      default: "Image upload"
    }
  },
  data() {
    return {
      file: {}
    }
  }
}
</script>
