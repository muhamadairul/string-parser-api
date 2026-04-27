"use client";

import { useState } from "react";
import InputForm from "@/components/InputForm";
import ResultCard from "@/components/ResultCard";
import HistoryTable from "@/components/HistoryTable";

interface ParseResult {
  id: number;
  name: string;
  age: string;
  city: string;
  raw: { name: string; age: string; city: string };
  created_at: string;
}

export default function HomePage() {
  const [result, setResult] = useState<ParseResult | null>(null);
  const [historyTrigger, setHistoryTrigger] = useState(0);

  const handleResult = (data: ParseResult) => {
    setResult(data);
  };

  const handleHistoryRefresh = () => {
    setHistoryTrigger((prev) => prev + 1);
  };

  return (
    <main className="main-layout">
      {/* Hero Header */}
      <div className="hero-section">
        <div className="hero-badge">🇮🇩 Indonesian String Parser</div>
        <h1 className="hero-title">
          <span className="gradient-text">String Parser API</span>
        </h1>
        <p className="hero-desc">
          Parsing otomatis <strong>Nama · Umur · Kota</strong> dari satu string input.
          Didukung deteksi ibukota provinsi &amp; output fixed-width.
        </p>
        <div className="hero-chips">
          <span className="chip chip-go">Go + Fiber</span>
          <span className="chip chip-next">Next.js 14</span>
          <span className="chip chip-pg">PostgreSQL</span>
          <span className="chip chip-algo">No Regex · RTL Parse</span>
        </div>
      </div>

      {/* Main Grid */}
      <div className="content-grid">
        {/* Left: Input + Result */}
        <div className="left-col">
          <InputForm onResult={handleResult} onHistoryRefresh={handleHistoryRefresh} />
          <ResultCard result={result} />
        </div>

        {/* Right: History */}
        <div className="right-col">
          <HistoryTable refreshTrigger={historyTrigger} />

          {/* Info Card */}
          <div className="info-card">
            <h4 className="info-title">📌 Constraint Parsing</h4>
            <ul className="info-list">
              <li><span className="info-icon">↩</span> Kanan ke kiri (RTL)</li>
              <li><span className="info-icon">🚫</span> No regex, no replace</li>
              <li><span className="info-icon">5️⃣</span> Max 5 variabel aktif</li>
              <li><span className="info-icon">📏</span> Name=30, Age=3, City=20 char</li>
              <li><span className="info-icon">🏙️</span> Deteksi 31 ibukota provinsi</li>
            </ul>
            <div className="info-suffix">
              <p className="info-suffix-title">Age suffix yang didukung:</p>
              <div className="suffix-chips">
                <code>28</code>
                <code>28TH</code>
                <code>28THN</code>
                <code>28 TAHUN</code>
                <code>28TAHUN</code>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
}
