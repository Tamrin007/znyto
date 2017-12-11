<template>
  <div class="siimple-box siimple-box--white">
    <div class="siimple-box-title">
      Converter
    </div>
    <div class="siimple-grid-row form">
      <div class="siimple-grid-col siimple-grid-col--5">
        <div class="siimple-form">
          <div class="siimple-form-title">
            ZNY
          </div>
          <div class="siimple-form-field">
            <input v-model="zny" type="number" step=any class="siimple-input siimple-input--fluid">
          </div>
        </div>
      </div>
      <div class="siimple-grid-col siimple-grid-col--2">
        <i class="fa fa-chevron-right fa-2x" aria-hidden="true"></i>
      </div>
      <div class="siimple-grid-col siimple-grid-col--5">
        <div class="siimple-form">
          <div class="siimple-form-title">
            JPY
          </div>
          <div class="siimple-form-field">
            <input v-model="jpy" type="number" step=any class="siimple-input siimple-input--fluid">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Converter',
  props: {
    rate: Number,
  },
  data() {
    return {
      zny: 1,
    };
  },
  computed: {
    jpy: {
      get() {
        return this.zny * this.rate;
      },
    },
  },
};
</script>

<style scoped>
.form {
  display: flex;
  align-items: center;
  text-align: center;
}

input {
  width: 90%;
  height: 60px;
  font-size: 2rem;
}
</style>
