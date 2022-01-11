<script lang="ts">
    import { onDestroy } from "svelte";
    import { API_BASE } from "../config";
    import { userStore } from "../userSession";

    const userMatcher = {
        '2': '1',
        '1': '2'
    }
    const user = $userStore

    let messages: any[] = []

    const getMessages = async () => {
        const url = `${API_BASE}/api/messages?`+ new URLSearchParams({
            recipient_id: userMatcher[user.id],
            amount: '100',
            offset: '0'
        })

        const resp = await fetch(url, {headers: {'session': user.authToken}})
        const json = await resp.json()

        messages = json
    }

    const interval = setInterval(() => getMessages(), 1000)

    onDestroy(() => clearInterval(interval))

    let sendMessageContent: string

    const sendMessage = async () => {
        const resp = await fetch(API_BASE+'/api/messages', {
            method: 'POST',
            headers: {
                'content-type': 'application/json',
                'session': user.authToken
            },
            body: JSON.stringify({
                content: sendMessageContent,
                recipient_id: userMatcher[user.id]
            })
        })

        console.log(resp)
    }
</script>

<div id="wrapper">
    <ul class="message-list">
        {#each messages as message}
            <li>{message.content}</li>
        {/each}
    </ul>

    <div class="send-wrapper">
        <input type="text" id="send-message-content" bind:value={sendMessageContent}>
        <input type="button" value="send message" id="send-message-button" on:click={() => sendMessage()}>
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
</style>
