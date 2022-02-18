import Home from "./routes/Home.svelte"
import Login from './routes/Login.svelte'
import Landing from './routes/Landing.svelte'
import Chat from './routes/Chat.svelte'

export default {
    '/': Home,
    '/login': Login,
    '/landing': Landing,
    '/chat/:id': Chat
}