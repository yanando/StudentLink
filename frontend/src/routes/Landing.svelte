<script lang="ts">
  import Loading from "../components/Loading.svelte";

  import { push } from "svelte-spa-router";
  import { API_BASE } from "../config";
  import { userStore } from "../userSession";

  const user = $userStore;

  let users: any[] = [];

  const getUsers = async () => {
    const url = `${API_BASE}/api/users`;
    const resp = await fetch(url, { headers: { session: user.authToken } });
    const json = await resp.json();

    return json;
  };

  getUsers().then((newUsers) => {
    users = newUsers;
  });
</script>

<div id="wrapper">
  <div class="inside-wrapper">
    <div id="chat-header">Users</div>
    <ul class="user-list">
      {#if users.length === 0}
        <Loading />
      {/if}
      {#each users.filter((allUsers) => allUsers.username !== user.username) as user}
        <li class="user" on:click={() => push(`/chat/${user.id}`)}>
          {user.username}
          <span>{user.type}</span>
        </li>
      {/each}
    </ul>
  </div>
</div>

<style lang="scss">
  #wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;

    font-size: 1em;
    background-color: #252427;
  }

  .inside-wrapper {
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    align-items: center;
    height: 80vh;
    min-width: 300px;
    width: 25vw;

    color: #7dafb9;
    background-color: #1f1f1e;
  }

  ul {
    display: flex;
    flex-direction: column;
    align-items: center;
    list-style: none;
    width: 80%;
    height: 80%;
    padding-inline-start: 0;
    overflow-y: scroll;

    &::-webkit-scrollbar {
      width: 1px;
    }
  }

  #chat-header {
    display: flex;
    justify-content: center;
    align-items: center;
    font-weight: 600;
  }

  .user {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
    font-size: 1.25em;
    height: 4em;
    margin: 10px 0;
    background-color: rgba(255, 255, 255, 0.1);
    border-radius: 10px;
    cursor: pointer;

    span {
      margin-right: 20px;
      align-self: flex-end;
      color: #527177;
    }
  }
</style>
