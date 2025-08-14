import React, { memo, useCallback, useState } from "react";

/* memo
useMemo
useCallback


*/

function HooksTuto() {
  const [parent, setParent] = useState(0);
  const [child1, setC] = useState(0);
  const [child2, setChild2] = useState(0);
  console.log("parent rendered");

  const setChild1 = useCallback(() => {
    setC((pre) => pre + 1);
  }, []);
  return (
    <div>
      <h2>parent - {parent}</h2>

      <button
        className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:ring-2 focus:ring-blue-400 focus:outline-none"
        onClick={() => setParent((pre) => pre + 1)}
      >
        Incre parent
      </button>

      <Child1 child1={child1} setChild1={setChild1} />

      <Child2 child2={child2} setChild2={setChild2} />
    </div>
  );
}

export default HooksTuto;

const Child1 = memo(({ child1, setChild1 }) => {
  console.log("child 1 rendered");
  return (
    <div>
      <h3>child 1 = {child1}</h3>
      <button
        className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:ring-2 focus:ring-blue-400 focus:outline-none"
        onClick={setChild1}
      >
        Incre Child1
      </button>
    </div>
  );
});

const Child2 = memo(({ child2, setChild2 }) => {
  console.log("child 2 rendered");

  return (
    <div>
      <h2>child-2 {child2}</h2>
      <button
        className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:ring-2 focus:ring-blue-400 focus:outline-none"
        onClick={() => setChild2((pre) => pre + 1)}
      >
        Incre Child2
      </button>
    </div>
  );
});
