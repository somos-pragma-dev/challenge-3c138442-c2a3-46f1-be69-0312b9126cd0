# Prompt para Mejorar el Codigo Base

Copia y pega el siguiente contenido completo en un asistente de IA (Claude, ChatGPT, etc.)
para obtener un ZIP con el proyecto corregido y listo para compilar.

---

```
Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Exportación en ZIP
Empaqueta todos los archivos en un único archivo ZIP descargable respetando exactamente
la estructura de rutas indicada por los marcadores.
El ZIP debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: El ZIP debe estar listo para descargar al finalizar. No preguntes si el usuario
quiere generarlo. Simplemente genera el archivo y proporciona el enlace de descarga; No debes desplegar en el chat el resumen de lo que arreglaste al Zip, solo entregalo.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación o descripciones sin código, genera los archivos
correspondientes sin aplicar análisis de compilación
No agregues texto después del enlace de descarga del ZIP
No preguntes si el usuario quiere el ZIP: simplemente generalo siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:
// === ARCHIVO: cmd/main.go ===
package main

import (
	"github.com/gin-gonic/gin"
	"internal/app/handlers"
)

func main() {
	r := gin.Default()
	accountHandler := handlers.NewAccountHandler()
	accountHandler.RegisterRoutes(r)
	r.Run(":8080")
}

// === ARCHIVO: internal/app/handlers/account_handler.go ===
package handlers

import (
	"github.com/gin-gonic/gin"
	"internal/app/models"
	"internal/app/services"
	"net/http"
)

type AccountHandler struct {
	accountService *services.AccountService
}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{
		accountService: services.NewAccountService(),
	}
}

func (h *AccountHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/accounts", h.CreateAccount)
	r.GET("/accounts/:id", h.GetAccount)
	r.PUT("/accounts/:id", h.UpdateAccount)
	r.DELETE("/accounts/:id", h.DeleteAccount)
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.accountService.CreateAccount(&account); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, account)
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	account, err := h.accountService.GetAccount(id)
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.Account
	if err := c.ShouldBindJSON(&account); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.accountService.UpdateAccount(id, &account); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if err := h.accountService.DeleteAccount(id); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// === ARCHIVO: internal/app/models/account.go ===
package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Number   string `json:"number"`
	Balance  float64 `json:"balance"`
	Owner    string `json:"owner"`
}

// === ARCHIVO: internal/app/services/account_service.go ===
package services

import (
	"internal/app/models"
	"internal/app/repositories"
)

type AccountService struct {
	accountRepository *repositories.AccountRepository
}

func NewAccountService() *AccountService {
	return &AccountService{
		accountRepository: repositories.NewAccountRepository(),
	}
}

func (s *AccountService) CreateAccount(account *models.Account) error {
	return s.accountRepository.CreateAccount(account)
}

func (s *AccountService) GetAccount(id string) (*models.Account, error) {
	return s.accountRepository.GetAccount(id)
}

func (s *AccountService) UpdateAccount(id string, account *models.Account) error {
	return s.accountRepository.UpdateAccount(id, account)
}

func (s *AccountService) DeleteAccount(id string) error {
	return s.accountRepository.DeleteAccount(id)
}

// === ARCHIVO: internal/app/repositories/account_repository.go ===
package repositories

import (
	"gorm.io/gorm"
	"internal/app/models"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository() *AccountRepository {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) // Placeholder for actual database connection
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) CreateAccount(account *models.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) GetAccount(id string) (*models.Account, error) {
	var account models.Account
	if err := r.db.First(&account, id).Error; err!= nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) UpdateAccount(id string, account *models.Account) error {
	return r.db.Model(&models.Account{}).Where("id =?", id).Updates(account).Error
}

func (r *AccountRepository) DeleteAccount(id string) error {
	return r.db.Delete(&models.Account{}, id).Error
}

// === ARCHIVO: pkg/errors/errors.go ===
package errors

import (
	"errors"
	"net/http"
)

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewAppError(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

// === ARCHIVO: tests/account_handler_test.go ===
package tests

import (
	"internal/app/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	r := gin.Default()
	accountHandler := handlers.NewAccountHandler()
	accountHandler.RegisterRoutes(r)
	req, _ := http.NewRequest("POST", "/accounts", strings.NewReader(`{"number":"123456789","balance":1000.0,"owner":"John Doe"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code!= http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, w.Code)
	}
}

// === ARCHIVO: config/config.yaml ===
# Configuration for the application
database:
  driver: sqlite3
  path: test.db
```
