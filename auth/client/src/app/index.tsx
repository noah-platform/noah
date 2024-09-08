import React from "react";
import { Button } from "@noah/ui/button";

function App(): JSX.Element {
  return (
    <div className="my-4">
      <h1 className="text-lg">Hello, World</h1>
      <Button
        onClick={() => {
          // eslint-disable-next-line no-alert -- alert for demo
          alert("Hi");
        }}
      >
        Hello
      </Button>
    </div>
  );
}

export default App;
