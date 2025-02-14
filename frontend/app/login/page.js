"use client";

import { signIn, signOut, useSession } from "next-auth/react";

export default function LoginPage() {
  const { data: session } = useSession();

  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      {session ? (
        <div className="text-center">
          <h2 className="text-lg font-semibold">Halo, {session.user.name}!</h2>
          <p>Email: {session.user.email}</p>
          <button
            onClick={() => signOut()}
            className="mt-4 px-4 py-2 bg-red-600 text-white rounded-md"
          >
            Logout
          </button>
        </div>
      ) : (
        <button
          onClick={() => signIn("google")}
          className="px-4 py-2 bg-teal-600 text-white rounded-md"
        >
          Login dengan Google
        </button>
      )}
    </div>
  );
}
