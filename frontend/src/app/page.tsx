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

  return (
    <main className="page">
      <header className="page-header">
        <h1>String Parser</h1>
        <p>Format input: <code>NAMA UMUR KOTA</code></p>
      </header>

      <div className="page-body">
        <section className="left-col">
          <InputForm
            onResult={setResult}
            onHistoryRefresh={() => setHistoryTrigger((n) => n + 1)}
          />
          <ResultCard result={result} />
        </section>

        <section className="right-col">
          <HistoryTable refreshTrigger={historyTrigger} />
        </section>
      </div>
    </main>
  );
}