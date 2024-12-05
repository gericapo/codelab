# Advent of Code: Trebuchet Calibration Problem


### Input
A text file containing lines \( T = \{L_1, L_2, \dots, L_n\} \), where each line \( L_i \) consists of:
- Numeric characters (\( \Sigma_{\text{num}} \)): \( \{0, 1, \dots, 9\} \)
- Alphabetic characters (\( \Sigma_{\text{alpha}} \)): \( \{a, b, \dots, z\} \)
- Word representations of numbers (\( \Sigma_{\text{words}} \)): \( \{\text{one, two, three, \dots, nine}\} \)

### Find
For each line \( L_i \), compute a **calibration value** \( V_i \) by extracting:
1. The **first digit** \( d_f \)
2. The **last digit** \( d_l \)

Digits may appear as numerals (\( \Sigma_{\text{num}} \)) or words (\( \Sigma_{\text{words}} \)). Combine them into a two-digit number:
\[
V_i = 10 \cdot d_f + d_l
\]

Finally, compute the total sum:
\[
S = \sum_{i=1}^n V_i
\]

---

Define the alphabet \( \Sigma \):
\[
\Sigma = \Sigma_{\text{num}} \cup \Sigma_{\text{alpha}} \cup \Sigma_{\text{words}}
\]
where:
- \( \Sigma_{\text{num}} = \{0, 1, \dots, 9\} \)
- \( \Sigma_{\text{alpha}} = \{a, b, \dots, z\} \)
- \( \Sigma_{\text{words}} = \{\text{one, two, three, \dots, nine}\} \)

Define a **tokenizer** \( \tau: \Sigma^* \to \{\text{digit, word, other}\}^* \), which splits a string \( L_i \) into **tokens**:
- A **digit token** \( d \in \Sigma_{\text{num}} \) represents a numeral.
- A **word token** \( w \in \Sigma_{\text{words}} \) represents a word mapped to a digit.
- All other characters form an **other token**.

Define a function \( M: \Sigma_{\text{words}} \to \Sigma_{\text{num}} \):
\[
M(\text{one}) = 1, \, M(\text{two}) = 2, \, \dots, \, M(\text{nine}) = 9
\]

- Define the language of **numerals**:
  \[
  L_{\text{num}} = \{x \in \Sigma^* \mid x \in \Sigma_{\text{num}}\}
  \]
- Define the language of **number words**:
  \[
  L_{\text{words}} = \{x \in \Sigma^* \mid x \in \Sigma_{\text{words}}\}
  \]
- Define the language of valid lines:
  \[
  L_{\text{valid}} = (\Sigma_{\text{num}} \cup \Sigma_{\text{words}} \cup \Sigma_{\text{alpha}})^*
  \]

---


For a given \( L_i \in T \):
1. Tokenize \( L_i \) into a sequence \( \tau(L_i) = [t_1, t_2, \dots, t_k] \), where \( t_j \) is either a digit, word, or other token.
2. Identify the **first and last tokens** that belong to \( L_{\text{num}} \cup L_{\text{words}} \):
   - First token: \( d_f = \min \{j \mid t_j \in L_{\text{num}} \cup L_{\text{words}}\} \)
   - Last token: \( d_l = \max \{j \mid t_j \in L_{\text{num}} \cup L_{\text{words}}\} \)

For tokens \( d_f, d_l \):
- If \( d_f, d_l \in \Sigma_{\text{num}} \), retain them.
- If \( d_f, d_l \in \Sigma_{\text{words}} \), map them using \( M \).

Combine \( d_f \) and \( d_l \) into a two-digit number:
\[
V_i = 10 \cdot d_f + d_l
\]

---



