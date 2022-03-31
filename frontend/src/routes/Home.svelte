<script lang="ts">
  import { push } from "svelte-spa-router";

  import Loading from "../components/Loading.svelte";
  setTimeout(() => {
    push("/login");
  }, 1000);
</script>

<div id="wrapper">
  <Loading />
</div>

<style lang="scss">
  #wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100vw;
    height: 100vh;
    background-color: #252427;
  }
</style>
