"use server"

import { revalidatePath } from "next/cache";
import { signIn, signOut } from "./auth";
import { Router } from "next/router";

export const login = async (prevState, formData) => {

    const { email, password } = Object.fromEntries(formData);

    try {
        await signIn("credentials", { email, password });
    } catch (err) {
    }
};

export const logOut = async () => {
    await signOut();
};