<script lang="ts">
  import { push } from "svelte-spa-router";
</script>

<div id="wrapper">
  <div class="inside-wrapper">
    <p class="back-button" on:click={() => push("/login")}>{"<"}</p>
    <h1>StudentLink Register</h1>
    <div>
      <p>
        Om dit systeem te kunnen gebruiken krijg je een account van je school,
        als deze school samenwerkt met StudentLink.
      </p>

      <p>Je hoeft dus niet zelf een account te registreren.</p>
    </div>
  </div>
</div>

<style lang="scss">
  #wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;

    font-size: 1.5vmin;
    background-color: #252427;
  }

  .inside-wrapper {
    position: relative;

    display: flex;
    flex-direction: column;
    align-items: center;
    height: 50vh;
    width: 25vw;
    min-width: 300px;

    color: #7dafb9;
    background-color: #1f1f1e;
  }

  .inside-wrapper h1 {
    margin: 10%;
  }

  .inside-wrapper div {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 20%;
    width: 80%;
  }

  @media (max-height: 900px) {
    .inside-wrapper div {
      font-size: 2em !important;
    }
  }

  .back-button {
    position: absolute;
    font-size: 1.5em;
    top: 10px;
    left: 17.5px;
    cursor: pointer;
  }
</style>
