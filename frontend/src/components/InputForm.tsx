"use client";

import { useState } from "react";

interface Props {
  onResult: (data: {
    id: number;
    name: string;
    age: string;
    city: string;
    raw: { name: string; age: string; city: string };
    created_at: string;
  }) => void;
  onHistoryRefresh: () => void;
}

const EXAMPLES = [
  "CUT MINI 28 BANDA ACEH",
  "BUDI SANTOSO 35THN SURABAYA",
  "SITI RAHAYU 22TH JAKARTA",
  "AHMAD YANI 40 TAHUN MAKASSAR",
];

export default function InputForm({ onResult, onHistoryRefresh }: Props) {
  const [input, setInput] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleParse = async () => {
    if (!input.trim()) {
      setError("Input tidak boleh kosong.");
      return;
    }
    setError("");
    setLoading(true);
    try {
      const res = await fetch("http://localhost:8080/api/parse", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ input: input.trim() }),
      });
      const json = await res.json();
      if (!res.ok || json.status !== "success") {
        throw new Error(json.message || "Gagal parsing");
      }
      onResult(json.data);
      onHistoryRefresh();
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : "Terjadi kesalahan");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="card">
      <label className="label" htmlFor="parse-input">Input string</label>
      <div className="input-row">
        <input
          id="parse-input"
          type="text"
          className={`text-input${error ? " text-input--error" : ""}`}
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={(e) => e.key === "Enter" && handleParse()}
          placeholder="Contoh: CUT MINI 28 BANDA ACEH"
          autoComplete="off"
          spellCheck={false}
        />
        <button className="btn-primary" onClick={handleParse} disabled={loading}>
          {loading ? "..." : "Parse"}
        </button>
      </div>

      {error && <p className="error-text">{error}</p>}

      <div className="examples">
        <span className="examples-label">Contoh:</span>
        {EXAMPLES.map((ex) => (
          <button key={ex} className="chip" onClick={() => setInput(ex)}>
            {ex}
          </button>
        ))}
      </div>
    </div>
  );
}