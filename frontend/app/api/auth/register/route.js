import { NextResponse } from "next/server";
import { createClient } from "@supabase/supabase-js";

const supabase = createClient(
  process.env.NEXT_PUBLIC_SUPABASE_URL,
  process.env.SUPABASE_SERVICE_ROLE_KEY
);

export async function POST(req) {
  try {
    const { id, name, email, avatar_url } = await req.json();

    // Cek apakah user sudah ada di database berdasarkan email
    const { data: existingUser, error: fetchError } = await supabase
      .from("players")
      .select("id")
      .eq("email", email)
      .single();

    if (fetchError && fetchError.code !== "PGRST116") {
      return NextResponse.json({ error: fetchError.message }, { status: 500 });
    }

    if (!existingUser) {
      // Jika user belum ada, tambahkan ke database
      const { error: insertError } = await supabase.from("players").insert([
        {
          id: id || undefined, // Gunakan id dari Google jika tersedia
          name,
          email,
          avatar_url,
        },
      ]);

      if (insertError) {
        return NextResponse.json(
          { error: insertError.message },
          { status: 500 }
        );
      }
    }

    return NextResponse.json(
      { message: "User registered successfully" },
      { status: 200 }
    );
  } catch (error) {
    return NextResponse.json({ error: error.message }, { status: 500 });
  }
}
