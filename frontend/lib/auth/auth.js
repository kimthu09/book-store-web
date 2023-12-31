import NextAuth from 'next-auth'
import CredentialsProvider from 'next-auth/providers/credentials'
import { authConfig } from './auth.config';
import axios from 'axios';

const login = async (credentials) => {
    try {
        const response = await axios.post("http://localhost:8080" + '/v1/login', {
            "email": credentials.email,
            "password": credentials.password
        })
        const { accessToken, refreshToken } = response.data.data;

        const user = await fetch("http://localhost:8080/v1/profile", {
            headers: {
                accept: "application/json",
                Authorization: `Bearer ${accessToken.token}`,
            },
        }).then((res) => res.json());
        user.token = accessToken.token
        return user

    } catch (e) {
        console.log(e.message)
    }

};

export const {
    handlers: { GET, POST },
    auth,
    signIn,
    signOut
} = NextAuth({
    ...authConfig,
    providers: [
        CredentialsProvider({
            async authorize(credentials, request) {
                try {
                    const user = await login(credentials);
                    return user;
                } catch (err) {
                    return null;
                }
            },
        }),
    ],
    callbacks: {
        ...authConfig.callbacks,
    }
})