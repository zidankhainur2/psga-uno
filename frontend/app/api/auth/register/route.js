import { NextResponse } from "next/server";
import { createClient } from "@supabase/supabase-js";

const supabase = createClient(
  process.env.NEXT_PUBLIC_SUPABASE_URL,
  process.env.SUPABASE_SERVICE_ROLE_KEY
);

export async function POST(req) {
  try {
    const { id, name, email, avatar_url } = await req.json();
    console.log("📩 Data yang diterima dari NextAuth:", {
      id,
      name,
      email,
      avatar_url,
    });

    // Cek apakah user sudah ada di database
    const { data: existingUser, error: fetchError } = await supabase
      .from("players")
      .select("*")
      .eq("email", email)
      .maybeSingle();

    console.log("🔍 Cek user di database:", existingUser, fetchError);

    if (fetchError && fetchError.code !== "PGRST116") {
      console.error("❌ Error saat fetch user:", fetchError);
      return NextResponse.json({ error: fetchError.message }, { status: 500 });
    }

    if (!existingUser) {
      console.log("🆕 User tidak ditemukan, menambahkan ke database...");

      // Jika user belum ada, tambahkan ke database
      const { error: insertError } = await supabase.from("players").insert([
        {
          name: name,
          email: email,
          avatar_url: avatar_url,
        },
      ]);

      if (insertError) {
        console.error("❌ Error saat insert user:", insertError);
        return NextResponse.json(
          { error: insertError.message },
          { status: 500 }
        );
      }

      console.log("✅ User berhasil ditambahkan ke database!");
    } else {
      console.log("✅ User sudah ada di database, tidak perlu insert.");
    }

    return NextResponse.json(
      { message: "User registered successfully" },
      { status: 200 }
    );
  } catch (error) {
    console.error("❌ Error di handler register:", error);
    return NextResponse.json({ error: error.message }, { status: 500 });
  }
}
