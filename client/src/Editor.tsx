import { VFC, useRef, useState, useEffect } from "react";
import * as monaco from "monaco-editor/esm/vs/editor/editor.api";
import styles from "./Editor.module.css";

export const Editor: VFC = () => {
  const [editor, setEditor] =
    useState<monaco.editor.IStandaloneCodeEditor | null>(null);
  const monacoEl = useRef(null);
  const [value, setValue] = useState<string>(
    "console.log(){  name = 'Regis',and ssome shiiit  }s",
  );

  useEffect(() => {
    if (monacoEl) {
      setEditor((editor) => {
        if (editor) return editor;
        const newEditor = monaco.editor.create(monacoEl.current!, {
          value: [
            "function x() {",
            '\tconsole.log("Hello world!");',
            "}",
          ].join(),
          language: "typescript",
        });

        newEditor.onDidChangeModelContent(() => {
          setValue(newEditor.getValue());
        });

        return newEditor;
      });
    }

    return () => editor?.dispose();
  }, [monacoEl.current]);

  console.log("Value:", JSON.stringify(value).split("\\n").join(","));

  return <div className={styles.Editor} ref={monacoEl}></div>;
};
