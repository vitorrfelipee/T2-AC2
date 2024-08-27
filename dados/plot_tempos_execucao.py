import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

# Carregar o CSV
df = pd.read_csv('tempos_execucao.csv')

# Configurar o estilo do gráfico
sns.set(style="whitegrid")

# Criar uma figura para os gráficos
plt.figure(figsize=(14, 8))

# Operações para serem plotadas
operacoes = df['Operacao'].unique()

# Plotar gráficos para cada operação
for i, operacao in enumerate(operacoes, 1):
    plt.subplot(2, 2, i)
    
    # Filtrar dados por operação
    df_operacao = df[df['Operacao'] == operacao]
    
    # Plotar dados sequenciais e paralelos para cada tamanho
    sns.lineplot(x='Tamanho', y='Tempo_medio', hue='Tipo', data=df_operacao, marker='o')
    
    # Configurar título e labels
    plt.title(f'Tempo de Execução: {operacao}')
    plt.xlabel('Tamanho da Matriz')
    plt.ylabel('Tempo Médio (segundos)')
    plt.yscale('log')  # Usar escala logarítmica para melhor visualização
    plt.grid(True)

# Ajustar o layout dos subplots
plt.tight_layout()

# Mostrar o gráfico
plt.show()
