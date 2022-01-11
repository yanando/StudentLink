import { writable } from 'svelte/store'

export interface user {
    id: number;
    type: 'university_student' | 'highschool_student';

    email: string;
    username: string;

    firstname: string;
    lastname: string;

    authToken: string;
}

export const userStore = writable<user>(JSON.parse(localStorage.getItem('user')) || {})
userStore.subscribe(e => localStorage.setItem('user', JSON.stringify(e)))
