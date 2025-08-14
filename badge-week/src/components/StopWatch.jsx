import { useEffect, useState } from "react";

function StopWatch() {
  const [seconds, setSeconds] = useState(0);
  const [running, setRunning] = useState(false);

  useEffect(() => {
    let timer;
    if (running) {
      timer = setInterval(() => {
        setSeconds((pre) => pre + 1);
      }, 1000);
    }

    return () => {
      console.log("intervel cleared");
      clearInterval(timer);
    };
  }, [running]);

  const start = () => {
    setRunning(true);
  };

  const stop = () => {
    setRunning(false);
  };

  const reset = () => {
    setRunning(false);
    setSeconds(0);
  };
  return (
    <div>
      <h2>{seconds}</h2>
      {running ? (
        <button onClick={stop}>Pause</button>
      ) : (
        <button onClick={start}>Start</button>
      )}
      <button onClick={reset}>Reset</button>
    </div>
  );
}

export default StopWatch;
