<template>
  <div id="app">
    <NavBar/>
    <ZnyTo/>
    <PageFooter/>
  </div>
</template>

<script>
import NavBar from './components/globals/NavBar';
import ZnyTo from './components/ZnyTo/ZnyTo';
import PageFooter from './components/globals/PageFooter';

export default {
  name: 'app',
  components: {
    NavBar,
    ZnyTo,
    PageFooter,
  },
};
</script>
