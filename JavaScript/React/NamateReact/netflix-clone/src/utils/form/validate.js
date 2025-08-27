export const checkValidation = (email, password) => {
  const isEmailValid = true; // .test(email)
  const isPasswordValid = true; // test(password)
  if (!isEmailValid) return "Email ID is not valid";
  if (!isPasswordValid) return "Password is not valid";

  return null;
};
