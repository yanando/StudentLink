import { writable } from 'svelte/store'

export interface user {
    id: number;
    type: 'university_student' | 'highschool_student';

    email: string;
    username: string;

    firstname: string;
    lastname: string;
}

export const userStore = writable<user>()