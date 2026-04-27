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

export default function InputForm({ onResult, onHistoryRefresh }: Props) {
  const [input, setInput] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleParse = async () => {
    if (!input.trim()) {
      setError("Input tidak boleh kosong!");
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
      const message = err instanceof Error ? err.message : "Terjadi kesalahan";
      setError(message);
    } finally {
      setLoading(false);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      handleParse();
    }
  };

  const examples = [
    "CUT MINI 28 BANDA ACEH",
    "BUDI SANTOSO 35THN SURABAYA",
    "SITI RAHAYU 22TH JAKARTA",
    "AHMAD YANI 40 TAHUN MAKASSAR",
  ];

  return (
    <div className="input-form-card">
      <div className="form-header">
        <div className="form-icon">⚡</div>
        <div>
          <h2 className="form-title">String Parser</h2>
          <p className="form-subtitle">Masukkan string dengan format: NAMA UMUR KOTA</p>
        </div>
      </div>

      <div className="input-group">
        <input
          id="parse-input"
          type="text"
          className={`parse-input ${error ? "input-error" : ""}`}
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={handleKeyDown}
          placeholder="Contoh: CUT MINI 28 BANDA ACEH"
          autoComplete="off"
          spellCheck={false}
        />
        <button
          id="parse-btn"
          className={`parse-btn ${loading ? "loading" : ""}`}
          onClick={handleParse}
          disabled={loading}
        >
          {loading ? (
            <span className="spinner" />
          ) : (
            <>
              <span>Parse</span>
              <span className="btn-arrow">→</span>
            </>
          )}
        </button>
      </div>

      {error && <p className="error-msg">⚠ {error}</p>}

      <div className="examples-section">
        <p className="examples-label">Coba contoh:</p>
        <div className="examples-list">
          {examples.map((ex) => (
            <button
              key={ex}
              className="example-chip"
              onClick={() => setInput(ex)}
            >
              {ex}
            </button>
          ))}
        </div>
      </div>
    </div>
  );
}
