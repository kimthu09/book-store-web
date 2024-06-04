import { NextResponse } from "next/server";

export const authConfig = {
  pages: {
    signIn: "/login",
    forgotPassword: "/forgot-password",
  },
  providers: [],
  callbacks: {
    jwt({ token, user }) {
      // console.log(user)
      if (user) {
        token.id = user;
      }
      return token;
    },
    session({ session, token }) {
      if (token) {
        session.user = token.id;
      }
      return session;
    },
    async authorized({ auth, request }) {
      const user = auth?.user;
      // console.log(user)
      const isOnLoginPage = request.nextUrl?.pathname.startsWith("/login");
      const isOnForgotPasswordPage =
        request.nextUrl?.pathname.startsWith("/forgot-password");

      if (user && isOnLoginPage) {
        return NextResponse.redirect(new URL("/", request.nextUrl));
      }
      if (user && isOnForgotPasswordPage) {
        return NextResponse.redirect(new URL("/", request.nextUrl));
      }
      if (!user && !isOnLoginPage && !isOnForgotPasswordPage) {
        return false;
      }
      return true;
    },
  },
};
