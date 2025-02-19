"use client";

import { useState, useEffect, Suspense } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import {
  Table,
  TableHead,
  TableRow,
  TableHeader,
  TableBody,
  TableCell,
} from "@/components/ui/table";
import { Dialog, DialogTrigger, DialogContent } from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";

export default function PlayersPage() {
  const [players, setPlayers] = useState(null); // Awalnya null agar tidak render langsung
  const [newPlayer, setNewPlayer] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    fetchPlayers();
  }, []);

  async function fetchPlayers() {
    try {
      const res = await fetch("http://localhost:8080/players/");
      const data = await res.json();
      setPlayers(data);
    } catch (error) {
      console.error("Error fetching players:", error);
    }
  }

  async function addPlayer() {
    if (!newPlayer) return;
    setLoading(true);
    try {
      await fetch("http://localhost:8080/players", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name: newPlayer }),
      });
      setNewPlayer("");
      fetchPlayers();
    } catch (error) {
      console.error("Error adding player:", error);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Manajemen Pemain</h1>
      <Dialog>
        <DialogTrigger asChild>
          <Button variant="default">Tambah Pemain</Button>
        </DialogTrigger>
        <DialogContent>
          <h2 className="text-lg font-bold">Tambah Pemain</h2>
          <Input
            value={newPlayer}
            onChange={(e) => setNewPlayer(e.target.value)}
            placeholder="Nama pemain"
          />
          <Button onClick={addPlayer} disabled={loading}>
            {loading ? "Menambah..." : "Tambah"}
          </Button>
        </DialogContent>
      </Dialog>
      <Card className="mt-4">
        <CardContent>
          <Suspense fallback={<p className="text-center">Loading...</p>}>
            {players === null ? (
              <p className="text-center">Mengambil data...</p>
            ) : (
              <Table>
                <TableHead>
                  <TableRow>
                    <TableHeader>No</TableHeader>
                    <TableHeader>Nama</TableHeader>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {players.map((player, index) => (
                    <TableRow key={player.id}>
                      <TableCell>{index + 1}</TableCell>
                      <TableCell>{player.name}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            )}
          </Suspense>
        </CardContent>
      </Card>
    </div>
  );
}
