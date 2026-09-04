# Implementación de una API REST con Go y Gin

El equipo de desarrollo de una fintech necesita implementar una API REST para gestionar operaciones de cuentas bancarias. La API debe permitir crear, leer, actualizar y eliminar cuentas. Además, debe manejar correctamente los errores y validaciones del dominio, como cuentas con saldo insuficiente o intentos de eliminación de cuentas con transacciones pendientes.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | Go Gin API |
| **Nivel** | junior-l1 |
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

### Fase 1: Definición del modelo de datos

**Objetivo:** Definir el modelo de datos para las cuentas bancarias, incluyendo atributos y relaciones.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Identificar los atributos necesarios para una cuenta bancaria (número de cuenta, saldo, propietario, etc.)
- Definir las relaciones entre las cuentas y otras entidades del dominio (transacciones, usuarios, etc.)
- Establecer las restricciones y validaciones necesarias para el modelo de datos

**Entregable:** Modelo de datos definido y documentado.

<details>
<summary>Pistas de conocimiento</summary>

- Considerar las relaciones entre las cuentas y las transacciones
- Evaluar las restricciones de negocio para las cuentas bancarias

</details>

### Fase 2: Implementación de las rutas de la API

**Objetivo:** Implementar las rutas de la API para crear, leer, actualizar y eliminar cuentas bancarias.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Crear las rutas para las operaciones CRUD de las cuentas bancarias
- Implementar la lógica de negocio para cada ruta, incluyendo validaciones y manejo de errores
- Asegurar que la API maneje correctamente los errores del dominio, como saldo insuficiente o transacciones pendientes

**Entregable:** Rutas de la API implementadas y funcionando.

<details>
<summary>Pistas de conocimiento</summary>

- Utilizar Gin para definir las rutas y los manejadores de las solicitudes
- Implementar la lógica de negocio en los manejadores de las rutas

</details>

### Fase 3: Pruebas y validación de la API

**Objetivo:** Realizar pruebas unitarias y de integración para validar el funcionamiento de la API.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Escribir pruebas unitarias para los manejadores de las rutas
- Realizar pruebas de integración para validar el flujo completo de las operaciones CRUD
- Asegurar que la API maneje correctamente los casos de error y las validaciones del dominio

**Entregable:** Pruebas unitarias y de integración implementadas y ejecutadas.

<details>
<summary>Pistas de conocimiento</summary>

- Utilizar un framework de pruebas para Go para escribir las pruebas
- Implementar pruebas de integración que cubran los casos de éxito y error

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es una cuenta bancaria y cuáles son sus atributos principales?
- **paraQueSirve**: ¿Para qué sirve la API REST en el contexto de una fintech?
- **comoSeUsa**: ¿Cómo se utilizan las rutas de la API para realizar operaciones CRUD en las cuentas bancarias?
- **erroresComunes**: ¿Cuáles son los errores comunes que puede encontrar la API al manejar cuentas bancarias y cómo se manejan?

## Criterios de Evaluacion

- Definición correcta del modelo de datos para las cuentas bancarias.
- Implementación de las rutas de la API para las operaciones CRUD.
- Manejo correcto de los errores y validaciones del dominio en la API.
- Pruebas unitarias y de integración implementadas y ejecutadas.

---

*Reto generado automaticamente por Challenge Generator - Pragma*
