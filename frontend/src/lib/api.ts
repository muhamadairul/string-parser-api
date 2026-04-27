const API_BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export interface ParseResult {
  id: number;
  name: string;
  age: string;
  city: string;
  raw: {
    name: string;
    age: string;
    city: string;
  };
  created_at: string;
}

export interface HistoryItem {
  id: number;
  name: string;
  age: string;
  city: string;
  created_at: string;
}

export async function parseString(input: string): Promise<{
  status: string;
  message: string;
  data: ParseResult;
}> {
  const res = await fetch(`${API_BASE}/api/parse`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ input }),
  });
  const json = await res.json();
  if (!res.ok) throw new Error(json.message || "Gagal parsing");
  return json;
}

export async function getHistory(): Promise<{
  status: string;
  message: string;
  data: HistoryItem[];
}> {
  const res = await fetch(`${API_BASE}/api/history`, {
    cache: "no-store",
  });
  const json = await res.json();
  if (!res.ok) throw new Error(json.message || "Gagal mengambil history");
  return json;
}
