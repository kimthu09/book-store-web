import { NextResponse } from "next/server";

export const authConfig = {
  pages: {
    signIn: "/login",
  },
  providers: [],
  callbacks: {
    jwt({ token, user }) {
      if (user) {
        token.id = user;
      }
      return token;
    },
    session({ session, token }) {
      if (token) {
        const tokenExpiresAt = new Date(
          new Date(token.id.token.created).getTime() +
            token.id.token.expiry * 1000
        );
        const currentTime = new Date();
        if (currentTime > tokenExpiresAt) {
          session.user = null;
          return session;
        }
        session.user = token.id;
      }
      return session;
    },
    async authorized({ auth, request }) {
      const user = auth?.user;

      const isOnLoginPage = request.nextUrl?.pathname.startsWith("/login");
      const isOnForgotPasswordPage =
        request.nextUrl?.pathname.startsWith("/forgot-password");

      if (user && (isOnLoginPage || isOnForgotPasswordPage)) {
        return NextResponse.redirect(new URL("/", request.nextUrl));
      }

      if (!user && !isOnLoginPage && !isOnForgotPasswordPage) {
        return false;
      }
      return true;
    },
  },
};
