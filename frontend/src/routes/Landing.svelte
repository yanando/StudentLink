<script lang="ts">
  import { onDestroy } from "svelte";
  import { API_BASE } from "../config";
  import { userStore } from "../userSession";
  import moment from "moment";

  const userMatcher = {
    "2": "1",
    "1": "3",
    "3": "1",
  };
  const user = $userStore;
  let recipientName = "";

  let messages: any[] = [];

  let hasScrolled = false;

  const scrollToEnd = () => {
    document
      .getElementsByClassName("message-list")[0]
      .scrollTo(
        0,
        document.getElementsByClassName("message-list")[0].scrollHeight
      );
  };

  const getMessages = async () => {
    const url =
      `${API_BASE}/api/messages?` +
      new URLSearchParams({
        recipient_id: userMatcher[user.id],
        amount: "100",
        offset: "0",
      });

    const resp = await fetch(url, { headers: { session: user.authToken } });
    const json = await resp.json();

    messages = json;
  };

  const getUser = async (id) => {
    const url = `${API_BASE}/api/user/${id}`;
    const resp = await fetch(url, { headers: { session: user.authToken } });
    const json = await resp.json();

    return json;
  };

  const interval = setInterval(async () => {
    const lastMessagesCount = messages.length;
    await getMessages();

    if (messages.length !== lastMessagesCount) {
      scrollToEnd();
      return;
    }

    if (!hasScrolled) {
      scrollToEnd();

      hasScrolled = true;
    }
  }, 1000);

  onDestroy(() => clearInterval(interval));

  let sendMessageContent: string;

  const sendMessage = async () => {
    await fetch(API_BASE + "/api/messages", {
      method: "POST",
      headers: {
        "content-type": "application/json",
        session: user.authToken,
      },
      body: JSON.stringify({
        content: sendMessageContent,
        recipient_id: parseInt(userMatcher[user.id]),
      }),
    });

    sendMessageContent = "";

    await getMessages();
    document
      .getElementsByClassName("message-list")[0]
      .scrollTo(
        0,
        document.getElementsByClassName("message-list")[0].scrollHeight
      );
  };

  getUser(userMatcher[user.id]).then((user) => {
    recipientName = user.username;
  });
</script>

<div id="wrapper">
  <div class="inside-wrapper">
    <div id="chat-header">
      <img src="https://via.placeholder.com/90x90.png" alt="icon" />
      {recipientName}
    </div>
    <ul class="message-list">
      {#each messages as message}
        <li class={message.author_id !== user.id ? "recipient-message" : ""}>
          <p>{message.content}</p><span>{moment(message.created_date).fromNow()}</span>
        </li>
      {/each}
    </ul>

    <div class="send-wrapper">
      <input
        type="text"
        id="send-message-content"
        bind:value={sendMessageContent}
        on:keyup={(event) =>
          event.key === "Enter" && event.ctrlKey ? sendMessage() : ""}
      />
      <input
        type="button"
        value="send message"
        id="send-message-button"
        on:click={() => sendMessage()}
      />
    </div>
  </div>
</div>

<style lang="scss">
  * {
    margin: 0;
  }

  body {
    overflow: hidden;
  }

  #wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;

    font-size: 1.5vmin;
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

  .send-wrapper {
    display: flex;
    justify-content: space-evenly;
    align-items: center;
  }

  ul {
    display: flex;
    flex-direction: column;
    list-style: none;
    width: 80%;
    height: 80%;
    padding-inline-start: 0;
    overflow-y: scroll;

    &::-webkit-scrollbar {
      width: 1px;
    }
  }

  li {
    position: relative;
    display: flex;
    justify-content: space-between;
    word-break: break-word;
    list-style: none;
    width: 100%;
    padding: 10px;

    &:hover {
      transition: 0.3s ease all;
      background-color: #181817;
    }

    /* Date */
    span {
      position: absolute;
      right: 5px;
      top: 5px;
      font-size: 1vmin;
    }
  }

  #send-message-content {
    width: 60%;
    text-align: center;
    outline: none;
    border: none;

    &::placeholder {
      color: white;
    }

    color: white;
    padding: 5px;
    font-size: 3vmin;
    border-radius: 5px;
    background-color: #252427;
  }

  #send-message-button {
    width: 30%;
    height: 100%;
    outline: none;

    font-size: 1.5vmin;
    border: none;
    border-radius: 5px;
    color: white;
    background-color: #0c7d9d;

    &:hover {
      cursor: pointer;
      transition: 0.3s ease all;
      background-color: rgba(12, 125, 157, 0.75);
    }
  }

  .recipient-message {
    color: yellow;
  }

  #chat-header {
    display: flex;
    justify-content: center;
    align-items: center;
    font-weight: 600;
    font-size: 2.5vmin;

    img {
      width: 30px;
      height: 30px;
      border-radius: 50%;
      margin-right: 20px;
    }
  }
</style>
