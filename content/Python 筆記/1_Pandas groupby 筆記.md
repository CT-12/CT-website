---
create_at: 2025.7.10
update_at: 2025.7.10
draft: false
tags: 
 - Python
 - Pandas
---

# Groupby

假設現在有這些資料：

|student|subject|semester|score|
|---|---|---|---|
|Alice|Math|2024-Fall|90|
|Alice|Science|2024-Fall|85|
|Bob|Math|2024-Fall|78|
|Bob|Science|2024-Fall|82|
|Alice|Math|2025-Spring|92|
|Bob|Math|2025-Spring|80|

然後對這些資料進行 `groupby`

```python
grouped_df = df.groupby(["student", "semester"])
```

接下來使用迴圈遍歷來看看 groupby 後資料長什麼樣子

```python
for (student, semester), group in grouped_df:
    print(f"Group: student={student}, semester={semester}")
    print()
    print(group)
    print()

```

## 第一次迴圈輸出

```python
Group: student=Alice, semester=2024-Fall

  student  subject    semester  score
0   Alice     Math  2024-Fall     90
1   Alice  Science  2024-Fall     85
```

## 第二次迴圈：
```python
Group: student=Alice, semester=2025-Spring

  student subject     semester  score
4   Alice    Math  2025-Spring     92
```

## 第三次迴圈：

```python
Group: student=Bob, semester=2024-Fall

  student  subject    semester  score
2     Bob     Math  2024-Fall     78
3     Bob  Science  2024-Fall     82
```
## 第四次迴圈：

```python
Group: student=Bob, semester=2025-Spring

  student subject     semester  score
5     Bob    Math  2025-Spring     80
```

使用 itertuples() 看單筆資料

```python
for (student, semester), group in grouped_df:
    print(f"Scores for {student} in {semester}:")
    for row in group.itertuples(index=False):
        print(f"  {row.subject}: {row.score}")
```

輸出會是：

```yaml
Scores for Alice in 2024-Fall:
  Math: 90
  Science: 85
Scores for Alice in 2025-Spring:
  Math: 92
...
```