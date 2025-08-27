const Logo = ({ size = "medium", showText = false }) => {
  // Size presets
  const sizes = {
    small: {
      container: { width: 120, height: 36 },
      text: { fontSize: 16 },
    },
    medium: {
      container: { width: 160, height: 48 },
      text: { fontSize: 24 },
    },
    large: {
      container: { width: 240, height: 72 },
      text: { fontSize: 36 },
    },
    xlarge: {
      container: { width: 320, height: 96 },
      text: { fontSize: 48 },
    },
  };

  const selectedSize = sizes[size] || sizes.medium;

  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        gap: 12,
      }}
    >
      {/* Red blocks part */}
      {/* <div style={{
        display: 'flex',
        width: selectedSize.container.width * 0.6,
        height: selectedSize.container.height
      }}>
        <div style={{
          width: '33.33%',
          height: '100%',
          backgroundColor: '#E50914'
        }}></div>
        <div style={{
          width: '33.33%',
          height: '100%',
          marginLeft: 4,
          backgroundColor: '#E50914',
          opacity: 0.9
        }}></div>
        <div style={{
          width: '33.33%',
          height: '100%',
          marginLeft: 4,
          backgroundColor: '#E50914',
          opacity: 0.8
        }}></div>
      </div> */}

      {/* Optional text part */}
      {showText && (
        <span
          style={{
            color: "#E50914",
            fontSize: selectedSize.text.fontSize,
            fontWeight: "bold",
            fontStyle: "italic",
            fontFamily: "Helvetica, Arial, sans-serif",
          }}
        >
          NETFLIX
        </span>
      )}
    </div>
  );
};

export default Logo;
