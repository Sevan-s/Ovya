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
        console.log('Suggestions :', results); // ← check ici

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

  onSubmit() {
    if (this.form.valid) {
      const value = this.form.value;

      value.date_start += ':00Z';
      value.date_end += ':00Z';

      this.http.post('http://localhost:8080/api/v1/visite/create', value)
        .subscribe({
          next: () => {
            alert('Visite créée avec succès !');
            this.form.reset();
          },
          error: (err) => {
            console.error(err);
            alert("Erreur lors de la création");
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
          },
          error: (err) => {
            console.error(err);
            alert('Erreur lors de la suppression');
          }
        });
    }
  }
}