# Desarrollo de una API REST para gestión de usuarios

El sistema de gestión de usuarios requiere una API REST que permita crear, leer, actualizar y eliminar usuarios. La API debe manejar solicitudes desde tres canales: web, móvil y API externa. Debe asegurar la idempotencia de las solicitudes y persistir cada solicitud con una clave única. Además, debe emitir un evento al sistema de auditoría por cada aceptación de solicitud.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | API REST en Go con Gin y GORM |
| **Nivel** | junior-l2 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Un IDE o editor de código.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Verifica que el proyecto arranca sin errores.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Implementación del canal de entrada

**Objetivo:** Crear un canal de entrada que acepte solicitudes y persista cada solicitud con idempotencia.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Identificar los tres canales de entrada y sus características.
- Definir la estructura de la solicitud y la respuesta.
- Implementar la persistencia de solicitudes con clave de idempotencia.

**Entregable:** Canal de entrada operativo persistiendo solicitudes con idempotencia.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo asegurar que dos solicitudes idénticas no generen duplicados.
- Piensa en cómo manejar diferentes formatos de solicitud desde los tres canales.

</details>

### Fase 2: Consolidación de canales y manejo de backpressure

**Objetivo:** Consolidar los tres canales de entrada sin pérdida de datos ante backpressure de uno de ellos.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Identificar posibles puntos de backpressure en cada canal.
- Implementar un mecanismo para manejar la backpressure sin perder solicitudes.
- Asegurar que la consolidación no afecte la idempotencia de las solicitudes.

**Entregable:** Consolidación de los tres canales sin pérdida ante backpressure de uno.

<details>
<summary>Pistas de conocimiento</summary>

- Considera el uso de colas o buffers para manejar la backpressure.
- Piensa en cómo mantener la idempotencia durante la consolidación.

</details>

### Fase 3: Recuperación automática ante caída del consumidor downstream

**Objetivo:** Implementar recuperación automática ante caída del consumidor downstream sin reprocesar solicitudes emitidas.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Identificar el consumidor downstream y sus posibles fallos.
- Implementar un mecanismo de recuperación automática.
- Asegurar que las solicitudes emitidas no se reprocesen en caso de caída del consumidor.

**Entregable:** Recuperación automática ante caída del consumidor downstream sin reprocesar emitidas.

<details>
<summary>Pistas de conocimiento</summary>

- Considera el uso de mecanismos de reintento con backoff exponencial.
- Piensa en cómo marcar las solicitudes emitidas para evitar reprocesamiento.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es la idempotencia y cómo se aplica en este contexto?
- **paraQueSirve**: ¿Para qué sirve la consolidación de canales en este sistema?
- **comoSeUsa**: ¿Cómo se usa un mecanismo de recuperación automática en este escenario?
- **erroresComunes**: ¿Cuáles son los errores comunes al manejar backpressure en este sistema?
- **queDecisionesImplica**: ¿Qué decisiones implica la implementación de un mecanismo de recuperación automática?

## Criterios de Evaluacion

- Implementación de un canal de entrada que acepte solicitudes y persista cada solicitud con idempotencia.
- Consolidación de los tres canales sin pérdida ante backpressure de uno.
- Recuperación automática ante caída del consumidor downstream sin reprocesar emitidas.

## Como trabajar con un asistente de IA

- **AGENTS.md** — instrucciones nativas del repo (Cursor, Codex, Copilot, Gemini, Claude Code). Abrí el proyecto y el agente las carga solo.
- **PROMPT_MEJORA.md** — el mismo prompt, para copiar y pegar en un chat (claude.ai, ChatGPT, etc.).

---

*Reto generado automaticamente por Challenge Generator - Pragma*
