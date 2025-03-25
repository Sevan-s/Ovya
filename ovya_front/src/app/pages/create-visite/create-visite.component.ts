import { Component, inject } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { switchMap } from 'rxjs/operators';
import { of } from 'rxjs';

interface Folder {
  Id: number;
  Ccial_id: number;
}

interface Acq {
  ID: number;
  Nom: string;
}

interface VisiteValue {
  date_start: string
  date_end: string
  acq_id: number | null
  ccial_id: number | null
  dossier_id: number | null
  canceled: boolean
}

@Component({
  selector: 'app-create-visite',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './create-visite.component.html',
  styleUrl: './create-visite.component.css'
})
export class CreateVisiteComponent {
  fb = inject(FormBuilder);
  http = inject(HttpClient);

  missingFields: string[] = [];
  form: FormGroup = this.fb.group({
    dossier_id: [''],
    date_start: [''],
    date_end: [''],
    acq_id: [''],
    acq_name: [''],
    ccial_id: [''],
    canceled: [false]
  });

  suggestions: Folder[] = [];
  suggestionsAcq: Acq[] = [];

  deleteAcqNameControl = new FormControl('');
  deleteSuggestions: Acq[] = [];
  allAcqs: Acq[] = [];
  idToDeleteAcqIsEmpty: boolean = false
  mode: 'create' | 'delete' = 'create';

  deleteControl = this.fb.control('');
  constructor() {
    this.deleteAcqNameControl.valueChanges.subscribe((name) => {
      if (name && this.allAcqs.length) {
        this.deleteSuggestions = this.allAcqs.filter(acq =>
          acq.Nom.toLowerCase().includes(name.toLowerCase())
        );
      } else {
        this.deleteSuggestions = [];
      }
    });
    this.http.get<Acq[]>('http://localhost:8080/api/v1/acq/all')
      .subscribe((data) => {
        this.allAcqs = data;
      });

    this.form.get('dossier_id')?.valueChanges
      .pipe(
        switchMap(value => {
          if (value) {
            return this.http.get<Folder[]>(`http://localhost:8080/api/v1/folder/search?query=${value}`);
          }
          return of([]);
        })
      )
      .subscribe((results) => {
        this.suggestions = results;
      });

    this.form.get('acq_name')?.valueChanges
      .subscribe((name) => {
        if (name && this.allAcqs.length) {
          this.suggestionsAcq = this.allAcqs.filter(acq =>
            acq.Nom.toLowerCase().includes(name.toLowerCase())
          );
        } else {
          this.suggestionsAcq = [];
        }
      });
  }


  selectSuggestion(folder: Folder) {
    this.form.patchValue({ dossier_id: folder.Id });
    this.suggestions = [];

    setTimeout(() => {
      this.suggestions = [];
    }, 100);
  }

  selectAcqSuggestion(acq: Acq) {
    this.form.patchValue({
      acq_name: acq.Nom,
      acq_id: acq.ID
    });
    this.suggestionsAcq = [];

    setTimeout(() => {
      this.suggestionsAcq = [];
    }, 100);
  }

  selectedAcqIdToDelete: number | null = null;

  selectAcqToDelete(acq: Acq) {
    this.deleteAcqNameControl.setValue(acq.Nom);
    this.selectedAcqIdToDelete = acq.ID;
    this.deleteSuggestions = [];
  }

  checkValueToCreateVisite(value: any): VisiteValue {
    const checkedValue: VisiteValue = {
      acq_id: value.acq_id ? parseInt(value.acq_id, 10) : null,
      canceled: value.canceled || false,
      ccial_id: value.ccial_id ? parseInt(value.ccial_id, 10) : null, 
      date_end: value.date_end || '',
      date_start: value.date_start || '',
      dossier_id: value.dossier_id ? parseInt(value.dossier_id, 10) : null,
    };
    
    console.log("Champs non vides : ", checkedValue);
    return checkedValue;
  }

  onSubmit() {
    if (this.form.valid) {
      const value = this.form.value;
      
      if (value.date_start) {
        value.date_start = new Date(value.date_start).toISOString();
      }
      
      if (value.date_end) {
        value.date_end = new Date(value.date_end).toISOString();
      }
      const cleanedValue: VisiteValue = this.checkValueToCreateVisite(value);
  
      this.missingFields = [];
      if (!cleanedValue['date_start']) {
        this.missingFields.push("Date de début");
      }
      if (!cleanedValue['dossier_id']) {
        this.missingFields.push("Id du dossier");
      }
      if (!cleanedValue['date_end']) {
        this.missingFields.push("Date de fin");
      }
      if (!cleanedValue['acq_id']) {
        this.missingFields.push("Nom de l'acquéreur");
      }
      if (!cleanedValue['ccial_id']) {
        this.missingFields.push("ID du commercial");
      }

      if (this.missingFields.length > 0) {
        return;
      }

  
      this.http.post('http://localhost:8080/api/v1/visite/create', cleanedValue)
      .subscribe({
        next: () => {
            console.log("value:", cleanedValue)
            alert('Visite créée avec succès !');
            this.form.reset();
            this.missingFields = [];
          },
          error: (err) => {
            console.error(err);
            if (err.status === 400) {
              console.log("value:", cleanedValue)
              console.log(err.error.fields)
              alert(`Erreur : ${err.error}`);
            } else {
              alert("Erreur lors de la création");
            }
          }
        });
    }
  }

  onDelete() {
    if (this.selectedAcqIdToDelete !== null) {
      this.http.delete(`http://localhost:8080/api/v1/acq/delete?id=${this.selectedAcqIdToDelete}`)
        .subscribe({
          next: () => {
            alert('Acquéreur supprimé');
            this.deleteAcqNameControl.reset();
            this.selectedAcqIdToDelete = null;
            this.idToDeleteAcqIsEmpty = false
          },
          error: (err) => {
            console.error(err);
            alert("Erreur lors de la suppression. Veuillez vérifié que l'acquéreur n'est pas lié à une visite.");
            this.idToDeleteAcqIsEmpty = false
          }
        });
    } else {
      this.idToDeleteAcqIsEmpty = true
    }
  }
}